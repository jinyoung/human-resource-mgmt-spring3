package human.resource.mgmt.domain;

import java.util.Date;
import java.util.List;
import javax.persistence.*;
import lombok.Data;

@Data
public class AddCalendarCommand {

    private String title;
    private Date from;
    private Date to;
}
