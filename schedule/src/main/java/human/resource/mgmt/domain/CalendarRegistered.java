package human.resource.mgmt.domain;

import human.resource.mgmt.domain.*;
import human.resource.mgmt.infra.AbstractEvent;
import java.util.*;
import lombok.*;

@Data
@ToString
public class CalendarRegistered extends AbstractEvent {

    private String userId;
    private String title;

    public CalendarRegistered(Calendar aggregate) {
        super(aggregate);
    }

    public CalendarRegistered() {
        super();
    }
}
