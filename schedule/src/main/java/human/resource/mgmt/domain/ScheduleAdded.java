package human.resource.mgmt.domain;

import human.resource.mgmt.domain.*;
import human.resource.mgmt.infra.AbstractEvent;
import java.util.*;
import lombok.*;

@Data
@ToString
public class ScheduleAdded extends AbstractEvent {

    private String userId;
    private String title;
    private Date date;

    public ScheduleAdded(Calendar aggregate) {
        super(aggregate);
    }

    public ScheduleAdded() {
        super();
    }
}
