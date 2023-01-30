package human.resource.mgmt.domain;

import human.resource.mgmt.domain.*;
import human.resource.mgmt.infra.AbstractEvent;
import java.util.*;
import lombok.*;

@Data
@ToString
public class ScheduleCanceled extends AbstractEvent {

    private String title;
    private Date date;
    private String userId;

    public ScheduleCanceled(Calendar aggregate) {
        super(aggregate);
    }

    public ScheduleCanceled() {
        super();
    }
}
